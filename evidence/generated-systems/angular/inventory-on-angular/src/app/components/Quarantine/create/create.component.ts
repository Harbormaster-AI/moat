import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QuarantineService } from '../../../services/Quarantine.service';
import { Quarantine } from '../../../models/Quarantine';
import { SubBaseComponent } from '../../Quarantine/sub.base.component';

@Component({
    selector: 'app-create-quarantine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQuarantineComponent extends SubBaseComponent implements OnInit {

    title = 'Add Quarantine';

    quarantineForm: FormGroup;
    quarantine: Quarantine;

    constructor( http: HttpClient,
        private quarantineService: QuarantineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition): void {
        this.quarantineService
        .addQuarantine(reason, startedAt, releasedAt, Warehouse, Items, Lot, SerialNumbers, Disposition)
            .subscribe(() => {
                this.router.navigate(['/indexQuarantine']);
            });
    }

    ngOnInit(): void {
    }
}