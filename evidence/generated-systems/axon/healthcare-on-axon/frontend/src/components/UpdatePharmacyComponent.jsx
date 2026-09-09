import React, { Component } from 'react'
import PharmacyService from '../services/PharmacyService';

class UpdatePharmacyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: ''
        }
        this.updatePharmacy = this.updatePharmacy.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        PharmacyService.getPharmacyById(this.state.id).then( (res) =>{
            let pharmacy = res.data;
            this.setState({
                name: pharmacy.name
            });
        });
    }

    updatePharmacy = (e) => {
        e.preventDefault();
        let pharmacy = {
            pharmacyId: this.state.id,
            name: this.state.name
        };
        console.log('pharmacy => ' + JSON.stringify(pharmacy));
        console.log('id => ' + JSON.stringify(this.state.id));
        PharmacyService.updatePharmacy(pharmacy).then( res => {
            this.props.history.push('/pharmacys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/pharmacys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Pharmacy</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePharmacy}>Save</button>
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

export default UpdatePharmacyComponent
