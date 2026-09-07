
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPurchaseOrderLineComponent } from './edit.component';
import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';

describe('EditPurchaseOrderLineComponent', () => {
  let component: EditPurchaseOrderLineComponent;
  let fixture: ComponentFixture<EditPurchaseOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPurchaseOrderLineComponent
      ],
      providers: [
        PurchaseOrderLineService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditPurchaseOrderLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});