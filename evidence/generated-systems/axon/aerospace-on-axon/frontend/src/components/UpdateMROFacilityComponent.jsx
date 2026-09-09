import React, { Component } from 'react'
import MROFacilityService from '../services/MROFacilityService';

class UpdateMROFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                approvalScope: '',
                address: ''
        }
        this.updateMROFacility = this.updateMROFacility.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeapprovalScopeHandler = this.changeapprovalScopeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
    }

    componentDidMount(){
        MROFacilityService.getMROFacilityById(this.state.id).then( (res) =>{
            let mROFacility = res.data;
            this.setState({
                name: mROFacility.name,
                approvalScope: mROFacility.approvalScope,
                address: mROFacility.address
            });
        });
    }

    updateMROFacility = (e) => {
        e.preventDefault();
        let mROFacility = {
            mROFacilityId: this.state.id,
            name: this.state.name,
            approvalScope: this.state.approvalScope,
            address: this.state.address
        };
        console.log('mROFacility => ' + JSON.stringify(mROFacility));
        console.log('id => ' + JSON.stringify(this.state.id));
        MROFacilityService.updateMROFacility(mROFacility).then( res => {
            this.props.history.push('/mROFacilitys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeapprovalScopeHandler= (event) => {
        this.setState({approvalScope: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }

    cancel(){
        this.props.history.push('/mROFacilitys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MROFacility</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> approvalScope: </label>
                                                <input placeholder="approvalScope" name="approvalScope" className="form-control" value={this.state.approvalScope} onChange={this.changeapprovalScopeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMROFacility}>Save</button>
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

export default UpdateMROFacilityComponent
