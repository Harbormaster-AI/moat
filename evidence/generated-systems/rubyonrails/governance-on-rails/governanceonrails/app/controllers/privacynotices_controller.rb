class PrivacyNoticesController < ApplicationController
  def index
    @privacyNotices = PrivacyNotice.all
  end
 
  def show
    @privacyNotice = PrivacyNotice.find(params[:id])
  end
 
  def new
    @privacyNotice = PrivacyNotice.new
  end
 
  def edit
    @privacyNotice = PrivacyNotice.find(params[:id])
  end
 
  def create
    @privacyNotice = PrivacyNotice.new(privacyNotice_params)
 
    if @privacyNotice.save
      redirect_to privacyNotices_path
    else
      render 'new'
    end
  end
 
  def update
    @privacyNotice = PrivacyNotice.find(params[:id])
 
    if @privacyNotice.update(privacyNotice_params)
      redirect_to privacyNotices_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @privacyNotice = PrivacyNotice.find(params[:id])
    @privacyNotice.destroy
    redirect_to privacyNotices_path
  end

 
  private
    def privacyNotice_params
      params.require(:privacyNotice).permit(:title, :audience, :versionLabel, :publicationDate, :publicationUrl, :Status)
    end
end