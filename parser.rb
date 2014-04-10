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
  rescue
    {}
  end

  def parse_amovies url
    doc = Nokogiri::HTML open url

    result = Hash[doc.xpath('//select[@id="series"]/option').map{|option| [option.content.force_encoding("utf-8"), option.attributes["value"].to_s]}]
    result[:title] = doc.css('.title_d_dot span').first.content
    result[:prefix] = doc.css('.prev_img img').attribute('title').text.strip
    result[:prefix] ||= result[:title][/[a-zA-Zа-яА-Я\s]+/].strip
    result
  end

  def get_series url
    serials_map = parse_amovies url
    Hash[serials_map.map do |episode, vk_link|
      if episode == :title
        [:title, vk_link]
      elsif episode == :prefix
        [:prefix, vk_link]
      else
        [episode, parse_vk(vk_link)]
      end
    end]
  end
  def russian_translit text 
    translited = text.tr('абвгдеёзийклмнопрстуфхэыь', 'abvgdeezijklmnoprstufhey\'')
    translited = translited.tr('АБВГДЕЁЗИЙКЛМНОПРСТУФХЭ', 'ABVGDEEZIJKLMNOPRSTUFHEY\'')

    translited = translited.gsub(/[жцчшщъюяЖЦЧШЩЪЮЯ]/,
        'ж' => 'zh', 'ц' => 'ts', 'ч' => 'ch', 'ш' => 'sh', 'щ' => 'sch', 'ъ' => '', 'ю' => 'ju', 'я' => 'ja',
        'Ж' => 'ZH', 'Ц' => 'TS', 'Ч' => 'CH', 'Ш' => 'SH', 'Щ' => 'SCH', 'Ъ' => '', 'Ю' => 'JU', 'Я' => 'JA')
    return translited
  end

  def get_filename prefix, episode, link
    [prefix,episode].map{|e| russian_translit(e)}.join(" ").gsub(" ", "_").downcase + ".mp4"
  end
  extend self
end
