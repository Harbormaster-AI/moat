
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTransferOrderLineComponent } from './create.component';
import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';
import { Router } from '@angular/router';

describe('CreateTransferOrderLineComponent', () => {
  let component: CreateTransferOrderLineComponent;
  let fixture: ComponentFixture<CreateTransferOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTransferOrderLineComponent
      ],
      providers: [
        TransferOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTransferOrderLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});