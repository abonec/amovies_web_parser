require 'nokogiri'
require 'open-uri'
module Parser
  def parse_vk url
    doc = Nokogiri::HTML open url

    value = doc.xpath('//object/embed').first.attributes["flashvars"].to_s
    value = URI.unescape value
    Hash[value.split("&").select{|param| param[/^url\d\d\d\=/]}.map do |param|
      param.match(/^url(\d\d\d)\=(.*)/).captures
    end]
  end

  def parse_amovies url
    doc = Nokogiri::HTML open url

    Hash[doc.xpath('//select[@id="series"]/option').map{|option| [option.content.force_encoding("utf-8"), option.attributes["value"].to_s]}]
  end

  def get_series url
    serials_map = parse_amovies url
    Hash[serials_map.map do |episode, vk_link|
      [episode, parse_vk(vk_link)]
    end]
  end
end
