import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QuarantineService } from '../../../services/Quarantine.service';
import { SubBaseComponent } from '../../Quarantine/sub.base.component';


@Component({
    selector: 'app-edit-quarantine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQuarantineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Quarantine';

    quarantineForm: FormGroup;
    quarantine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QuarantineService,
        private fb: FormBuilder
) {
        super(http);
        this.quarantineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reason: ['', Validators.required],
      startedAt: ['', Validators.required],
      releasedAt: ['', Validators.required],
      Warehouse: ['', ],
      Items: ['', ],
      Lot: ['', ],
      SerialNumbers: ['', ],
      Disposition: ['', ]
        });
    }

    
    updateQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQuarantine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQuarantine(params['id']).subscribe(res => {
                this.quarantine = res;
            });
        });
    }
}