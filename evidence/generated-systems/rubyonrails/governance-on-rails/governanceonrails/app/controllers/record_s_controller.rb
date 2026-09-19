class Record_sController < ApplicationController
  def index
    @record_s = Record_.all
  end
 
  def show
    @record_ = Record_.find(params[:id])
  end
 
  def new
    @record_ = Record_.new
  end
 
  def edit
    @record_ = Record_.find(params[:id])
  end
 
  def create
    @record_ = Record_.new(record__params)
 
    if @record_.save
      redirect_to record_s_path
    else
      render 'new'
    end
  end
 
  def update
    @record_ = Record_.find(params[:id])
 
    if @record_.update(record__params)
      redirect_to record_s_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @record_ = Record_.find(params[:id])
    @record_.destroy
    redirect_to record_s_path
  end

 
  private
    def record__params
      params.require(:record_).permit(:title, :creationDate, :RecordType, :Classification, :Status)
    end
end