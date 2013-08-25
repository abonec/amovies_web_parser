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
  uri = URI.parse "http://cs533322v4.vk.me/u219120229/videos/a26f0fd78c.720.mp4"
  req = Net::HTTP.new(uri.hostname, 80)
  req.request_head(uri.path).inspect
end
