class AirworthinessDirectivesController < ApplicationController
  def index
    @airworthinessDirectives = AirworthinessDirective.all
  end
 
  def show
    @airworthinessDirective = AirworthinessDirective.find(params[:id])
  end
 
  def new
    @airworthinessDirective = AirworthinessDirective.new
  end
 
  def edit
    @airworthinessDirective = AirworthinessDirective.find(params[:id])
  end
 
  def create
    @airworthinessDirective = AirworthinessDirective.new(airworthinessDirective_params)
 
    if @airworthinessDirective.save
      redirect_to airworthinessDirectives_path
    else
      render 'new'
    end
  end
 
  def update
    @airworthinessDirective = AirworthinessDirective.find(params[:id])
 
    if @airworthinessDirective.update(airworthinessDirective_params)
      redirect_to airworthinessDirectives_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @airworthinessDirective = AirworthinessDirective.find(params[:id])
    @airworthinessDirective.destroy
    redirect_to airworthinessDirectives_path
  end

 
  private
    def airworthinessDirective_params
      params.require(:airworthinessDirective).permit(:directiveNumber, :title)
    end
end