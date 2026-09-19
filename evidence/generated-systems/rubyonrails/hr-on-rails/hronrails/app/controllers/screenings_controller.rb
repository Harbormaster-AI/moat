class ScreeningsController < ApplicationController
  def index
    @screenings = Screening.all
  end
 
  def show
    @screening = Screening.find(params[:id])
  end
 
  def new
    @screening = Screening.new
  end
 
  def edit
    @screening = Screening.find(params[:id])
  end
 
  def create
    @screening = Screening.new(screening_params)
 
    if @screening.save
      redirect_to screenings_path
    else
      render 'new'
    end
  end
 
  def update
    @screening = Screening.find(params[:id])
 
    if @screening.update(screening_params)
      redirect_to screenings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @screening = Screening.find(params[:id])
    @screening.destroy
    redirect_to screenings_path
  end

 
  private
    def screening_params
      params.require(:screening).permit(:name, :completedDate, :Status)
    end
end