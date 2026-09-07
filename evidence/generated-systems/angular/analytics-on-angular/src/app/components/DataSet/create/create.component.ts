import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataSetService } from '../../../services/DataSet.service';
import { DataSet } from '../../../models/DataSet';
import { SubBaseComponent } from '../../DataSet/sub.base.component';

@Component({
    selector: 'app-create-dataSet',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataSetComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataSet';

    dataSetForm: FormGroup;
    dataSet: DataSet;

    constructor( http: HttpClient,
        private dataSetService: DataSetService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataSetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      schemaVersion: ['', Validators.required],
      refreshSchedule: ['', Validators.required],
      Sensitive: ['', Validators.required],
      Workspace: ['', ],
      Sources: ['', ],
      Pipelines: ['', ],
      SemanticModels: ['', ],
      Dimensions: ['', ],
      Measures: ['', ],
      Metrics: ['', ],
      QualityRules: ['', ],
      LineageNode: ['', ],
      Tags: ['', ],
      DataFormat: ['', ]
        });
    }

    
    addDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat): void {
        this.dataSetService
        .addDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat)
            .subscribe(() => {
                this.router.navigate(['/indexDataSet']);
            });
    }

    ngOnInit(): void {
    }
}