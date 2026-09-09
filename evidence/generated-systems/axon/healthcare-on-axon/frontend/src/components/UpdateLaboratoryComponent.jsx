import React, { Component } from 'react'
import LaboratoryService from '../services/LaboratoryService';

class UpdateLaboratoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                cliaNumber: ''
        }
        this.updateLaboratory = this.updateLaboratory.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecliaNumberHandler = this.changecliaNumberHandler.bind(this);
    }

    componentDidMount(){
        LaboratoryService.getLaboratoryById(this.state.id).then( (res) =>{
            let laboratory = res.data;
            this.setState({
                name: laboratory.name,
                cliaNumber: laboratory.cliaNumber
            });
        });
    }

    updateLaboratory = (e) => {
        e.preventDefault();
        let laboratory = {
            laboratoryId: this.state.id,
            name: this.state.name,
            cliaNumber: this.state.cliaNumber
        };
        console.log('laboratory => ' + JSON.stringify(laboratory));
        console.log('id => ' + JSON.stringify(this.state.id));
        LaboratoryService.updateLaboratory(laboratory).then( res => {
            this.props.history.push('/laboratorys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Laboratory</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> cliaNumber: </label>
                                                <input placeholder="cliaNumber" name="cliaNumber" className="form-control" value={this.state.cliaNumber} onChange={this.changecliaNumberHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLaboratory}>Save</button>
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

export default UpdateLaboratoryComponent
