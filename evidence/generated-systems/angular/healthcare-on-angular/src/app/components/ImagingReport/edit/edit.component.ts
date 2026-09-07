import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ImagingReportService } from '../../../services/ImagingReport.service';
import { SubBaseComponent } from '../../ImagingReport/sub.base.component';


@Component({
    selector: 'app-edit-imagingReport',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditImagingReportComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ImagingReport';

    imagingReportForm: FormGroup;
    imagingReport: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ImagingReportService,
        private fb: FormBuilder
) {
        super(http);
        this.imagingReportForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reportNumber: ['', Validators.required],
      impression: ['', Validators.required],
      reportedDate: ['', Validators.required],
      ImagingOrder: ['', ],
      Clinician: ['', ],
      Encounter: ['', ],
      ImagingCenter: ['', ],
      Status: ['', ]
        });
    }

    
    updateImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateImagingReport(reportNumber, impression, reportedDate, ImagingOrder, Clinician, Encounter, ImagingCenter, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexImagingReport']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getImagingReport(params['id']).subscribe(res => {
                this.imagingReport = res;
            });
        });
    }
}