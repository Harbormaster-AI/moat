class ControlTest_sController < ApplicationController
  def index
    @controlTest_s = ControlTest_.all
  end
 
  def show
    @controlTest_ = ControlTest_.find(params[:id])
  end
 
  def new
    @controlTest_ = ControlTest_.new
  end
 
  def edit
    @controlTest_ = ControlTest_.find(params[:id])
  end
 
  def create
    @controlTest_ = ControlTest_.new(controlTest__params)
 
    if @controlTest_.save
      redirect_to controlTest_s_path
    else
      render 'new'
    end
  end
 
  def update
    @controlTest_ = ControlTest_.find(params[:id])
 
    if @controlTest_.update(controlTest__params)
      redirect_to controlTest_s_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @controlTest_ = ControlTest_.find(params[:id])
    @controlTest_.destroy
    redirect_to controlTest_s_path
  end

 
  private
    def controlTest__params
      params.require(:controlTest_).permit(:name, :testPeriodStart, :testPeriodEnd, :sampleSize, :TestType, :Effectiveness, :Status)
    end
end