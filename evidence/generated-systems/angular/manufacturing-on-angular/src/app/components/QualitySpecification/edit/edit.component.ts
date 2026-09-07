import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QualitySpecificationService } from '../../../services/QualitySpecification.service';
import { SubBaseComponent } from '../../QualitySpecification/sub.base.component';


@Component({
    selector: 'app-edit-qualitySpecification',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQualitySpecificationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit QualitySpecification';

    qualitySpecificationForm: FormGroup;
    qualitySpecification: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QualitySpecificationService,
        private fb: FormBuilder
) {
        super(http);
        this.qualitySpecificationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  specCode: ['', Validators.required],
      name: ['', Validators.required],
      version: ['', Validators.required],
      Item: ['', ]
        });
    }

    
    updateQualitySpecification(specCode, name, version, Item): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQualitySpecification(specCode, name, version, Item, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQualitySpecification']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQualitySpecification(params['id']).subscribe(res => {
                this.qualitySpecification = res;
            });
        });
    }
}