import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QualityCheckService } from '../../../services/QualityCheck.service';
import { QualityCheck } from '../../../models/QualityCheck';
import { SubBaseComponent } from '../../QualityCheck/sub.base.component';

@Component({
    selector: 'app-create-qualityCheck',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQualityCheckComponent extends SubBaseComponent implements OnInit {

    title = 'Add QualityCheck';

    qualityCheckForm: FormGroup;
    qualityCheck: QualityCheck;

    constructor( http: HttpClient,
        private qualityCheckService: QualityCheckService,
        private fb: FormBuilder,
        private router: Router
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

    
    addQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status): void {
        this.qualityCheckService
        .addQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status)
            .subscribe(() => {
                this.router.navigate(['/indexQualityCheck']);
            });
    }

    ngOnInit(): void {
    }
}