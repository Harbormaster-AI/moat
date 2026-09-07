
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWorkOrderComponent } from './create.component';
import { WorkOrderService } from '../../../services/WorkOrder.service';
import { Router } from '@angular/router';

describe('CreateWorkOrderComponent', () => {
  let component: CreateWorkOrderComponent;
  let fixture: ComponentFixture<CreateWorkOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWorkOrderComponent
      ],
      providers: [
        WorkOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWorkOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});