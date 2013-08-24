require 'sinatra'
require 'haml'
require './parser'

include Parser

set :public_folder, 'public'

get '/' do
  haml :index
end

get '/links' do
  haml :links, locals: { series: get_series(params[:url]) }
end
