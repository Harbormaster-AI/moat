import React, { Component } from 'react'
import DSPService from '../services/DSPService';

class CreateDSPComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                website: '',
                region: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DSPService.getDSPById(this.state.id).then( (res) =>{
                let dSP = res.data;
                this.setState({
                    name: dSP.name,
                    website: dSP.website,
                    region: dSP.region
                });
            });
        }        
    }
    saveOrUpdateDSP = (e) => {
        e.preventDefault();
        let dSP = {
                dSPId: this.state.id,
                name: this.state.name,
                website: this.state.website,
                region: this.state.region
            };
        console.log('dSP => ' + JSON.stringify(dSP));

        // step 5
        if(this.state.id === '_add'){
            dSP.dSPId=''
            DSPService.createDSP(dSP).then(res =>{
                this.props.history.push('/dSPs');
            });
        }else{
            DSPService.updateDSP(dSP).then( res => {
                this.props.history.push('/dSPs');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeregionHandler= (event) => {
        this.setState({region: event.target.value});
    }

    cancel(){
        this.props.history.push('/dSPs');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DSP</h3>
        }else{
            return <h3 className="text-center">Update DSP</h3>
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

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> region:&emsp; </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDSP}>Save</button>
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

export default CreateDSPComponent
