import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BIQueryService } from '../../../services/BIQuery.service';
import { BIQuery } from '../../../models/BIQuery';
import { SubBaseComponent } from '../../BIQuery/sub.base.component';

@Component({
    selector: 'app-create-bIQuery',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBIQueryComponent extends SubBaseComponent implements OnInit {

    title = 'Add BIQuery';

    bIQueryForm: FormGroup;
    bIQuery: BIQuery;

    constructor( http: HttpClient,
        private bIQueryService: BIQueryService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.bIQueryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      text: ['', Validators.required],
      Workspace: ['', ],
      Datasets: ['', ],
      Reports: ['', ],
      Dashboards: ['', ],
      Notebooks: ['', ],
      Dialect: ['', ]
        });
    }

    
    addBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect): void {
        this.bIQueryService
        .addBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect)
            .subscribe(() => {
                this.router.navigate(['/indexBIQuery']);
            });
    }

    ngOnInit(): void {
    }
}