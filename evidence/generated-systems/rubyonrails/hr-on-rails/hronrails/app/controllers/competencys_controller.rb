class CompetencysController < ApplicationController
  def index
    @competencys = Competency.all
  end
 
  def show
    @competency = Competency.find(params[:id])
  end
 
  def new
    @competency = Competency.new
  end
 
  def edit
    @competency = Competency.find(params[:id])
  end
 
  def create
    @competency = Competency.new(competency_params)
 
    if @competency.save
      redirect_to competencys_path
    else
      render 'new'
    end
  end
 
  def update
    @competency = Competency.find(params[:id])
 
    if @competency.update(competency_params)
      redirect_to competencys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @competency = Competency.find(params[:id])
    @competency.destroy
    redirect_to competencys_path
  end

 
  private
    def competency_params
      params.require(:competency).permit(:name, :category)
    end
end