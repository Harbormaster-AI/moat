class WishlistItemsController < ApplicationController
  def index
    @wishlistItems = WishlistItem.all
  end
 
  def show
    @wishlistItem = WishlistItem.find(params[:id])
  end
 
  def new
    @wishlistItem = WishlistItem.new
  end
 
  def edit
    @wishlistItem = WishlistItem.find(params[:id])
  end
 
  def create
    @wishlistItem = WishlistItem.new(wishlistItem_params)
 
    if @wishlistItem.save
      redirect_to wishlistItems_path
    else
      render 'new'
    end
  end
 
  def update
    @wishlistItem = WishlistItem.find(params[:id])
 
    if @wishlistItem.update(wishlistItem_params)
      redirect_to wishlistItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @wishlistItem = WishlistItem.find(params[:id])
    @wishlistItem.destroy
    redirect_to wishlistItems_path
  end

 
  private
    def wishlistItem_params
      params.require(:wishlistItem).permit(:addedDate)
    end
end