class DataSubjectRequestsController < ApplicationController
  def index
    @dataSubjectRequests = DataSubjectRequest.all
  end
 
  def show
    @dataSubjectRequest = DataSubjectRequest.find(params[:id])
  end
 
  def new
    @dataSubjectRequest = DataSubjectRequest.new
  end
 
  def edit
    @dataSubjectRequest = DataSubjectRequest.find(params[:id])
  end
 
  def create
    @dataSubjectRequest = DataSubjectRequest.new(dataSubjectRequest_params)
 
    if @dataSubjectRequest.save
      redirect_to dataSubjectRequests_path
    else
      render 'new'
    end
  end
 
  def update
    @dataSubjectRequest = DataSubjectRequest.find(params[:id])
 
    if @dataSubjectRequest.update(dataSubjectRequest_params)
      redirect_to dataSubjectRequests_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataSubjectRequest = DataSubjectRequest.find(params[:id])
    @dataSubjectRequest.destroy
    redirect_to dataSubjectRequests_path
  end

 
  private
    def dataSubjectRequest_params
      params.require(:dataSubjectRequest).permit(:receivedDate, :dueDate, :requesterCountry, :RequestType, :Status)
    end
end