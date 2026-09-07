import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CreativeApprovalService } from '../../../services/CreativeApproval.service';
import { SubBaseComponent } from '../../CreativeApproval/sub.base.component';


@Component({
    selector: 'app-edit-creativeApproval',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCreativeApprovalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CreativeApproval';

    creativeApprovalForm: FormGroup;
    creativeApproval: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CreativeApprovalService,
        private fb: FormBuilder
) {
        super(http);
        this.creativeApprovalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reviewer: ['', Validators.required],
      reviewedAt: ['', Validators.required],
      CreativeAsset: ['', ],
      Publisher: ['', ],
      Status: ['', ]
        });
    }

    
    updateCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCreativeApproval']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCreativeApproval(params['id']).subscribe(res => {
                this.creativeApproval = res;
            });
        });
    }
}