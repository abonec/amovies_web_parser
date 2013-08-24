require 'nokogiri'
require 'open-uri'
require 'sinatra'
require 'haml'

def get_720p url
  doc = Nokogiri::HTML open url

  value = doc.xpath('//object/param[@name="flashvars"]').first.attributes["value"].to_s
  value = URI.unescape value
  value.split("&").select{|param| param[/^url\d\d\d\=/]}.map do |param|
    param.match(/^url\d\d\d\=(.*)/).captures.last
  end
end

# get_720p "http://vk.com/video_ext.php?oid=-56436575&id=165929277&hash=225039cd9b76e1e9&sd"


def get_series url
  doc = Nokogiri::HTML open url

  # value = doc.xpath('//select[@id="series"]/option').map{|option| option.attributes["value"].to_s}
  Hash[doc.xpath('//select[@id="series"]/option').map{|option| [option.content.force_encoding("utf-8"), option.attributes["value"].to_s]}]

end

puts get_series("http://amovies.tv/serials/490-vo-vse-tyazhkie.html")

# get_series("http://amovies.tv/serials/490-vo-vse-tyazhkie.html").each do |link|
#   puts get_720p(link)
# end


get '/' do
  "<a src='http://cs1-50v4.vk.me/p19/cded410d8122.720.mp4'>http://cs1-50v4.vk.me/p19/cded410d8122.720.mp4</a>"
end



__END__
@@ index
%html
 %head
  %body
    %form{action: "/links", method: :get}
      %fieldset
        %input{type: :text, name: :url}
      %input{ type: :submit, value: 'Get Serials'}

