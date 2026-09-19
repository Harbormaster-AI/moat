class ContactsController < ApplicationController
  def index
    @contacts = Contact.all
  end
 
  def show
    @contact = Contact.find(params[:id])
  end
 
  def new
    @contact = Contact.new
  end
 
  def edit
    @contact = Contact.find(params[:id])
  end
 
  def create
    @contact = Contact.new(contact_params)
 
    if @contact.save
      redirect_to contacts_path
    else
      render 'new'
    end
  end
 
  def update
    @contact = Contact.find(params[:id])
 
    if @contact.update(contact_params)
      redirect_to contacts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @contact = Contact.find(params[:id])
    @contact.destroy
    redirect_to contacts_path
  end

 
  private
    def contact_params
      params.require(:contact).permit(:firstName, :lastName, :title, :email, :phone, :mobile, :mailingAddress, :PreferredContactMethod)
    end
end