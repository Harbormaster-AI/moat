class MediaAssetsController < ApplicationController
  def index
    @mediaAssets = MediaAsset.all
  end
 
  def show
    @mediaAsset = MediaAsset.find(params[:id])
  end
 
  def new
    @mediaAsset = MediaAsset.new
  end
 
  def edit
    @mediaAsset = MediaAsset.find(params[:id])
  end
 
  def create
    @mediaAsset = MediaAsset.new(mediaAsset_params)
 
    if @mediaAsset.save
      redirect_to mediaAssets_path
    else
      render 'new'
    end
  end
 
  def update
    @mediaAsset = MediaAsset.find(params[:id])
 
    if @mediaAsset.update(mediaAsset_params)
      redirect_to mediaAssets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @mediaAsset = MediaAsset.find(params[:id])
    @mediaAsset.destroy
    redirect_to mediaAssets_path
  end

 
  private
    def mediaAsset_params
      params.require(:mediaAsset).permit(:url, :altText, :position, :MediaType)
    end
end