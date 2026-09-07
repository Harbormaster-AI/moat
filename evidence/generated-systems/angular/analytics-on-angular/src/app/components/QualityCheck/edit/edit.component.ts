import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QualityCheckService } from '../../../services/QualityCheck.service';
import { SubBaseComponent } from '../../QualityCheck/sub.base.component';


@Component({
    selector: 'app-edit-qualityCheck',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQualityCheckComponent extends SubBaseComponent implements OnInit {

    title = 'Edit QualityCheck';

    qualityCheckForm: FormGroup;
    qualityCheck: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QualityCheckService,
        private fb: FormBuilder
) {
        super(http);
        this.qualityCheckForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  checkedAt: ['', Validators.required],
      observedValue: ['', Validators.required],
      sampleSize: ['', Validators.required],
      Rule: ['', ],
      Dataset: ['', ],
      Status: ['', ]
        });
    }

    
    updateQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQualityCheck']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQualityCheck(params['id']).subscribe(res => {
                this.qualityCheck = res;
            });
        });
    }
}