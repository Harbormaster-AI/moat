
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCostCenterComponent } from './create.component';
import { CostCenterService } from '../../../services/CostCenter.service';
import { Router } from '@angular/router';

describe('CreateCostCenterComponent', () => {
  let component: CreateCostCenterComponent;
  let fixture: ComponentFixture<CreateCostCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCostCenterComponent
      ],
      providers: [
        CostCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCostCenterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});