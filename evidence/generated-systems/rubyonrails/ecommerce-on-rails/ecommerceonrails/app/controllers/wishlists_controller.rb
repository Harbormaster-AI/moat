class WishlistsController < ApplicationController
  def index
    @wishlists = Wishlist.all
  end
 
  def show
    @wishlist = Wishlist.find(params[:id])
  end
 
  def new
    @wishlist = Wishlist.new
  end
 
  def edit
    @wishlist = Wishlist.find(params[:id])
  end
 
  def create
    @wishlist = Wishlist.new(wishlist_params)
 
    if @wishlist.save
      redirect_to wishlists_path
    else
      render 'new'
    end
  end
 
  def update
    @wishlist = Wishlist.find(params[:id])
 
    if @wishlist.update(wishlist_params)
      redirect_to wishlists_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @wishlist = Wishlist.find(params[:id])
    @wishlist.destroy
    redirect_to wishlists_path
  end

 
  private
    def wishlist_params
      params.require(:wishlist).permit(:name, :asPublic, :createdAt)
    end
end