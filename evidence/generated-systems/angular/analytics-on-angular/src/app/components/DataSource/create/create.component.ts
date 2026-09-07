import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataSourceService } from '../../../services/DataSource.service';
import { DataSource } from '../../../models/DataSource';
import { SubBaseComponent } from '../../DataSource/sub.base.component';

@Component({
    selector: 'app-create-dataSource',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataSourceComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataSource';

    dataSourceForm: FormGroup;
    dataSource: DataSource;

    constructor( http: HttpClient,
        private dataSourceService: DataSourceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataSourceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      connection: ['', Validators.required],
      Streaming: ['', Validators.required],
      Workspace: ['', ],
      ProducedDatasets: ['', ],
      Pipelines: ['', ],
      SourceType: ['', ],
      Format: ['', ]
        });
    }

    
    addDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format): void {
        this.dataSourceService
        .addDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format)
            .subscribe(() => {
                this.router.navigate(['/indexDataSource']);
            });
    }

    ngOnInit(): void {
    }
}