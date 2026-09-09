import React, { Component } from 'react'
import InsurerService from '../services/InsurerService';

class CreateInsurerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                domicileCountry: '',
                naicNumber: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changedomicileCountryHandler = this.changedomicileCountryHandler.bind(this);
        this.changenaicNumberHandler = this.changenaicNumberHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InsurerService.getInsurerById(this.state.id).then( (res) =>{
                let insurer = res.data;
                this.setState({
                    name: insurer.name,
                    legalName: insurer.legalName,
                    domicileCountry: insurer.domicileCountry,
                    naicNumber: insurer.naicNumber,
                    website: insurer.website
                });
            });
        }        
    }
    saveOrUpdateInsurer = (e) => {
        e.preventDefault();
        let insurer = {
                insurerId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                domicileCountry: this.state.domicileCountry,
                naicNumber: this.state.naicNumber,
                website: this.state.website
            };
        console.log('insurer => ' + JSON.stringify(insurer));

        // step 5
        if(this.state.id === '_add'){
            insurer.insurerId=''
            InsurerService.createInsurer(insurer).then(res =>{
                this.props.history.push('/insurers');
            });
        }else{
            InsurerService.updateInsurer(insurer).then( res => {
                this.props.history.push('/insurers');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changedomicileCountryHandler= (event) => {
        this.setState({domicileCountry: event.target.value});
    }
    changenaicNumberHandler= (event) => {
        this.setState({naicNumber: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/insurers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Insurer</h3>
        }else{
            return <h3 className="text-center">Update Insurer</h3>
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

                                            <label> legalName:&emsp; </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> domicileCountry:&emsp; </label>
                                                <input placeholder="domicileCountry" name="domicileCountry" className="form-control" value={this.state.domicileCountry} onChange={this.changedomicileCountryHandler}/>

                                            <label> naicNumber:&emsp; </label>
                                                <input placeholder="naicNumber" name="naicNumber" className="form-control" value={this.state.naicNumber} onChange={this.changenaicNumberHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInsurer}>Save</button>
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

export default CreateInsurerComponent
