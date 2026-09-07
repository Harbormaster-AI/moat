
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditTransferOrderLineComponent } from './edit.component';
import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';

describe('EditTransferOrderLineComponent', () => {
  let component: EditTransferOrderLineComponent;
  let fixture: ComponentFixture<EditTransferOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditTransferOrderLineComponent
      ],
      providers: [
        TransferOrderLineService,
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

    fixture = TestBed.createComponent(EditTransferOrderLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});