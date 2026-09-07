
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateHealthSystemComponent } from './create.component';
import { HealthSystemService } from '../../../services/HealthSystem.service';
import { Router } from '@angular/router';

describe('CreateHealthSystemComponent', () => {
  let component: CreateHealthSystemComponent;
  let fixture: ComponentFixture<CreateHealthSystemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateHealthSystemComponent
      ],
      providers: [
        HealthSystemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateHealthSystemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});