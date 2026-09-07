
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePlannedOrderComponent } from './create.component';
import { PlannedOrderService } from '../../../services/PlannedOrder.service';
import { Router } from '@angular/router';

describe('CreatePlannedOrderComponent', () => {
  let component: CreatePlannedOrderComponent;
  let fixture: ComponentFixture<CreatePlannedOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePlannedOrderComponent
      ],
      providers: [
        PlannedOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePlannedOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});