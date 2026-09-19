class Exception_sController < ApplicationController
  def index
    @exception_s = Exception_.all
  end
 
  def show
    @exception_ = Exception_.find(params[:id])
  end
 
  def new
    @exception_ = Exception_.new
  end
 
  def edit
    @exception_ = Exception_.find(params[:id])
  end
 
  def create
    @exception_ = Exception_.new(exception__params)
 
    if @exception_.save
      redirect_to exception_s_path
    else
      render 'new'
    end
  end
 
  def update
    @exception_ = Exception_.find(params[:id])
 
    if @exception_.update(exception__params)
      redirect_to exception_s_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @exception_ = Exception_.find(params[:id])
    @exception_.destroy
    redirect_to exception_s_path
  end

 
  private
    def exception__params
      params.require(:exception_).permit(:title, :justification, :startDate, :endDate, :ExceptionType, :Status)
    end
end