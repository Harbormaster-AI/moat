
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEquityGrantComponent } from './create.component';
import { EquityGrantService } from '../../../services/EquityGrant.service';
import { Router } from '@angular/router';

describe('CreateEquityGrantComponent', () => {
  let component: CreateEquityGrantComponent;
  let fixture: ComponentFixture<CreateEquityGrantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEquityGrantComponent
      ],
      providers: [
        EquityGrantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEquityGrantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});