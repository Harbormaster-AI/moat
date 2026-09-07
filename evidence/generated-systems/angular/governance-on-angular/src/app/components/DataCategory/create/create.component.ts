import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataCategoryService } from '../../../services/DataCategory.service';
import { DataCategory } from '../../../models/DataCategory';
import { SubBaseComponent } from '../../DataCategory/sub.base.component';

@Component({
    selector: 'app-create-dataCategory',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataCategoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataCategory';

    dataCategoryForm: FormGroup;
    dataCategory: DataCategory;

    constructor( http: HttpClient,
        private dataCategoryService: DataCategoryService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataCategoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      description: ['', Validators.required],
      ProcessingActivities: ['', ],
      Records: ['', ],
      DataBreaches: ['', ],
      Classification: ['', ]
        });
    }

    
    addDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification): void {
        this.dataCategoryService
        .addDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification)
            .subscribe(() => {
                this.router.navigate(['/indexDataCategory']);
            });
    }

    ngOnInit(): void {
    }
}