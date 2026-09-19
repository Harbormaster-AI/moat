class EmailMessagesController < ApplicationController
  def index
    @emailMessages = EmailMessage.all
  end
 
  def show
    @emailMessage = EmailMessage.find(params[:id])
  end
 
  def new
    @emailMessage = EmailMessage.new
  end
 
  def edit
    @emailMessage = EmailMessage.find(params[:id])
  end
 
  def create
    @emailMessage = EmailMessage.new(emailMessage_params)
 
    if @emailMessage.save
      redirect_to emailMessages_path
    else
      render 'new'
    end
  end
 
  def update
    @emailMessage = EmailMessage.find(params[:id])
 
    if @emailMessage.update(emailMessage_params)
      redirect_to emailMessages_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @emailMessage = EmailMessage.find(params[:id])
    @emailMessage.destroy
    redirect_to emailMessages_path
  end

 
  private
    def emailMessage_params
      params.require(:emailMessage).permit(:subject, :body, :sentAt, :messageId, :Direction, :Status)
    end
end