require 'sinatra'
require 'haml'
require './parser'
require 'net/http'
require 'httparty'

include Parser
set :bind, '0.0.0.0'
set :port, '6100'

set :public_folder, 'public'

get '/' do
  haml :index
end

get '/links' do
  haml :links, locals: { series: get_series(params[:url]) }
end

post '/download' do
  prefix, episode, link = params[:prefix], params[:episode], params[:link]
  filename = get_filename prefix, episode, link
  HTTParty.get("http://localhost:4545/down?url=#{link}&filename=#{filename}").parsed_response
end
