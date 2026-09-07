
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRiskComponent } from './create.component';
import { RiskService } from '../../../services/Risk.service';
import { Router } from '@angular/router';

describe('CreateRiskComponent', () => {
  let component: CreateRiskComponent;
  let fixture: ComponentFixture<CreateRiskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRiskComponent
      ],
      providers: [
        RiskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRiskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});