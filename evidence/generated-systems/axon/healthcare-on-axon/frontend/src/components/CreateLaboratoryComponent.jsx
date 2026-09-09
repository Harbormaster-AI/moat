import React, { Component } from 'react'
import LaboratoryService from '../services/LaboratoryService';

class CreateLaboratoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                cliaNumber: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecliaNumberHandler = this.changecliaNumberHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LaboratoryService.getLaboratoryById(this.state.id).then( (res) =>{
                let laboratory = res.data;
                this.setState({
                    name: laboratory.name,
                    cliaNumber: laboratory.cliaNumber
                });
            });
        }        
    }
    saveOrUpdateLaboratory = (e) => {
        e.preventDefault();
        let laboratory = {
                laboratoryId: this.state.id,
                name: this.state.name,
                cliaNumber: this.state.cliaNumber
            };
        console.log('laboratory => ' + JSON.stringify(laboratory));

        // step 5
        if(this.state.id === '_add'){
            laboratory.laboratoryId=''
            LaboratoryService.createLaboratory(laboratory).then(res =>{
                this.props.history.push('/laboratorys');
            });
        }else{
            LaboratoryService.updateLaboratory(laboratory).then( res => {
                this.props.history.push('/laboratorys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecliaNumberHandler= (event) => {
        this.setState({cliaNumber: event.target.value});
    }

    cancel(){
        this.props.history.push('/laboratorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Laboratory</h3>
        }else{
            return <h3 className="text-center">Update Laboratory</h3>
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

                                            <label> cliaNumber:&emsp; </label>
                                                <input placeholder="cliaNumber" name="cliaNumber" className="form-control" value={this.state.cliaNumber} onChange={this.changecliaNumberHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLaboratory}>Save</button>
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

export default CreateLaboratoryComponent
