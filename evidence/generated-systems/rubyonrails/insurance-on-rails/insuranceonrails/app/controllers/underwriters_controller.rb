class UnderwritersController < ApplicationController
  def index
    @underwriters = Underwriter.all
  end
 
  def show
    @underwriter = Underwriter.find(params[:id])
  end
 
  def new
    @underwriter = Underwriter.new
  end
 
  def edit
    @underwriter = Underwriter.find(params[:id])
  end
 
  def create
    @underwriter = Underwriter.new(underwriter_params)
 
    if @underwriter.save
      redirect_to underwriters_path
    else
      render 'new'
    end
  end
 
  def update
    @underwriter = Underwriter.find(params[:id])
 
    if @underwriter.update(underwriter_params)
      redirect_to underwriters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @underwriter = Underwriter.find(params[:id])
    @underwriter.destroy
    redirect_to underwriters_path
  end

 
  private
    def underwriter_params
      params.require(:underwriter).permit(:firstName, :lastName, :employeeId, :authorityLimit)
    end
end