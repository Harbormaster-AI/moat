import React, { Component } from 'react'
import QualitySpecificationService from '../services/QualitySpecificationService';

class CreateQualitySpecificationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                specCode: '',
                name: '',
                version: ''
        }
        this.changespecCodeHandler = this.changespecCodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeversionHandler = this.changeversionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            QualitySpecificationService.getQualitySpecificationById(this.state.id).then( (res) =>{
                let qualitySpecification = res.data;
                this.setState({
                    specCode: qualitySpecification.specCode,
                    name: qualitySpecification.name,
                    version: qualitySpecification.version
                });
            });
        }        
    }
    saveOrUpdateQualitySpecification = (e) => {
        e.preventDefault();
        let qualitySpecification = {
                qualitySpecificationId: this.state.id,
                specCode: this.state.specCode,
                name: this.state.name,
                version: this.state.version
            };
        console.log('qualitySpecification => ' + JSON.stringify(qualitySpecification));

        // step 5
        if(this.state.id === '_add'){
            qualitySpecification.qualitySpecificationId=''
            QualitySpecificationService.createQualitySpecification(qualitySpecification).then(res =>{
                this.props.history.push('/qualitySpecifications');
            });
        }else{
            QualitySpecificationService.updateQualitySpecification(qualitySpecification).then( res => {
                this.props.history.push('/qualitySpecifications');
            });
        }
    }
    
    changespecCodeHandler= (event) => {
        this.setState({specCode: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeversionHandler= (event) => {
        this.setState({version: event.target.value});
    }

    cancel(){
        this.props.history.push('/qualitySpecifications');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add QualitySpecification</h3>
        }else{
            return <h3 className="text-center">Update QualitySpecification</h3>
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
                                            <label> specCode:&emsp; </label>
                                                <input placeholder="specCode" name="specCode" className="form-control" value={this.state.specCode} onChange={this.changespecCodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> version:&emsp; </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateQualitySpecification}>Save</button>
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

export default CreateQualitySpecificationComponent
