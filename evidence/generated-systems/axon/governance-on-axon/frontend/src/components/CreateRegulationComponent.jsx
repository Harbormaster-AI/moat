import React, { Component } from 'react'
import RegulationService from '../services/RegulationService';

class CreateRegulationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                citation: '',
                jurisdiction: '',
                publicationUrl: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecitationHandler = this.changecitationHandler.bind(this);
        this.changejurisdictionHandler = this.changejurisdictionHandler.bind(this);
        this.changepublicationUrlHandler = this.changepublicationUrlHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RegulationService.getRegulationById(this.state.id).then( (res) =>{
                let regulation = res.data;
                this.setState({
                    name: regulation.name,
                    citation: regulation.citation,
                    jurisdiction: regulation.jurisdiction,
                    publicationUrl: regulation.publicationUrl
                });
            });
        }        
    }
    saveOrUpdateRegulation = (e) => {
        e.preventDefault();
        let regulation = {
                regulationId: this.state.id,
                name: this.state.name,
                citation: this.state.citation,
                jurisdiction: this.state.jurisdiction,
                publicationUrl: this.state.publicationUrl
            };
        console.log('regulation => ' + JSON.stringify(regulation));

        // step 5
        if(this.state.id === '_add'){
            regulation.regulationId=''
            RegulationService.createRegulation(regulation).then(res =>{
                this.props.history.push('/regulations');
            });
        }else{
            RegulationService.updateRegulation(regulation).then( res => {
                this.props.history.push('/regulations');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecitationHandler= (event) => {
        this.setState({citation: event.target.value});
    }
    changejurisdictionHandler= (event) => {
        this.setState({jurisdiction: event.target.value});
    }
    changepublicationUrlHandler= (event) => {
        this.setState({publicationUrl: event.target.value});
    }

    cancel(){
        this.props.history.push('/regulations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Regulation</h3>
        }else{
            return <h3 className="text-center">Update Regulation</h3>
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

                                            <label> citation:&emsp; </label>
                                                <input placeholder="citation" name="citation" className="form-control" value={this.state.citation} onChange={this.changecitationHandler}/>

                                            <label> jurisdiction:&emsp; </label>
                                                <input placeholder="jurisdiction" name="jurisdiction" className="form-control" value={this.state.jurisdiction} onChange={this.changejurisdictionHandler}/>

                                            <label> publicationUrl:&emsp; </label>
                                                <input placeholder="publicationUrl" name="publicationUrl" className="form-control" value={this.state.publicationUrl} onChange={this.changepublicationUrlHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRegulation}>Save</button>
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

export default CreateRegulationComponent
