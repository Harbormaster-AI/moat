class RoutingsController < ApplicationController
  def index
    @routings = Routing.all
  end
 
  def show
    @routing = Routing.find(params[:id])
  end
 
  def new
    @routing = Routing.new
  end
 
  def edit
    @routing = Routing.find(params[:id])
  end
 
  def create
    @routing = Routing.new(routing_params)
 
    if @routing.save
      redirect_to routings_path
    else
      render 'new'
    end
  end
 
  def update
    @routing = Routing.find(params[:id])
 
    if @routing.update(routing_params)
      redirect_to routings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @routing = Routing.find(params[:id])
    @routing.destroy
    redirect_to routings_path
  end

 
  private
    def routing_params
      params.require(:routing).permit(:routingNumber, :revision, :effectivityStart, :effectivityEnd, :RoutingType, :Status)
    end
end