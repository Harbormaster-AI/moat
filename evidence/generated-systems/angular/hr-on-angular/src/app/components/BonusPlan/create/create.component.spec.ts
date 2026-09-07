
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBonusPlanComponent } from './create.component';
import { BonusPlanService } from '../../../services/BonusPlan.service';
import { Router } from '@angular/router';

describe('CreateBonusPlanComponent', () => {
  let component: CreateBonusPlanComponent;
  let fixture: ComponentFixture<CreateBonusPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBonusPlanComponent
      ],
      providers: [
        BonusPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBonusPlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});