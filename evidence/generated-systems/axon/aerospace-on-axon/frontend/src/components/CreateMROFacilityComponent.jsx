import React, { Component } from 'react'
import MROFacilityService from '../services/MROFacilityService';

class CreateMROFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                approvalScope: '',
                address: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeapprovalScopeHandler = this.changeapprovalScopeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MROFacilityService.getMROFacilityById(this.state.id).then( (res) =>{
                let mROFacility = res.data;
                this.setState({
                    name: mROFacility.name,
                    approvalScope: mROFacility.approvalScope,
                    address: mROFacility.address
                });
            });
        }        
    }
    saveOrUpdateMROFacility = (e) => {
        e.preventDefault();
        let mROFacility = {
                mROFacilityId: this.state.id,
                name: this.state.name,
                approvalScope: this.state.approvalScope,
                address: this.state.address
            };
        console.log('mROFacility => ' + JSON.stringify(mROFacility));

        // step 5
        if(this.state.id === '_add'){
            mROFacility.mROFacilityId=''
            MROFacilityService.createMROFacility(mROFacility).then(res =>{
                this.props.history.push('/mROFacilitys');
            });
        }else{
            MROFacilityService.updateMROFacility(mROFacility).then( res => {
                this.props.history.push('/mROFacilitys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MROFacility</h3>
        }else{
            return <h3 className="text-center">Update MROFacility</h3>
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

                                            <label> approvalScope:&emsp; </label>
                                                <input placeholder="approvalScope" name="approvalScope" className="form-control" value={this.state.approvalScope} onChange={this.changeapprovalScopeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMROFacility}>Save</button>
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

export default CreateMROFacilityComponent
