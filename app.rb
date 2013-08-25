require 'sinatra'
require 'haml'
require './parser'
require 'net/http'

include Parser

set :public_folder, 'public'

get '/' do
  haml :index
end

get '/links' do
  haml :links, locals: { series: get_series(params[:url]) }
end

require 'open-uri'
get '/heroku_test' do
  open 'http://vk.com/video_ext.php?oid=-56941034&id=165774609&hash=328e1a94b8d1bfff&sd'
end

get '/vk_head' do
  uri = URI.parse Parser.parse_vk("http://vk.com/video_ext.php?oid=-56436575&id=165929277&hash=225039cd9b76e1e9&sd").to_a.last.last
  req = Net::HTTP.new(uri.hostname, 80)
  req.request_head(uri.path).inspect
end
