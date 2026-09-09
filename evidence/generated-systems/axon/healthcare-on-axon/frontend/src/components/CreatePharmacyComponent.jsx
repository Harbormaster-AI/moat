import React, { Component } from 'react'
import PharmacyService from '../services/PharmacyService';

class CreatePharmacyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PharmacyService.getPharmacyById(this.state.id).then( (res) =>{
                let pharmacy = res.data;
                this.setState({
                    name: pharmacy.name
                });
            });
        }        
    }
    saveOrUpdatePharmacy = (e) => {
        e.preventDefault();
        let pharmacy = {
                pharmacyId: this.state.id,
                name: this.state.name
            };
        console.log('pharmacy => ' + JSON.stringify(pharmacy));

        // step 5
        if(this.state.id === '_add'){
            pharmacy.pharmacyId=''
            PharmacyService.createPharmacy(pharmacy).then(res =>{
                this.props.history.push('/pharmacys');
            });
        }else{
            PharmacyService.updatePharmacy(pharmacy).then( res => {
                this.props.history.push('/pharmacys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/pharmacys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Pharmacy</h3>
        }else{
            return <h3 className="text-center">Update Pharmacy</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePharmacy}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreatePharmacyComponent
