require 'sinatra'
require 'haml'
require './parser'

include Parser

get '/' do
  haml :index
end

get '/links' do
  get_series(params[:url]).inspect
end
