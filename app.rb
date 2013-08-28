require 'sinatra'
require 'haml'
require './parser'
require 'net/http'

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
