class CartItemsController < ApplicationController
  def index
    @cartItems = CartItem.all
  end
 
  def show
    @cartItem = CartItem.find(params[:id])
  end
 
  def new
    @cartItem = CartItem.new
  end
 
  def edit
    @cartItem = CartItem.find(params[:id])
  end
 
  def create
    @cartItem = CartItem.new(cartItem_params)
 
    if @cartItem.save
      redirect_to cartItems_path
    else
      render 'new'
    end
  end
 
  def update
    @cartItem = CartItem.find(params[:id])
 
    if @cartItem.update(cartItem_params)
      redirect_to cartItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @cartItem = CartItem.find(params[:id])
    @cartItem.destroy
    redirect_to cartItems_path
  end

 
  private
    def cartItem_params
      params.require(:cartItem).permit(:quantity, :unitPrice, :totalPrice)
    end
end