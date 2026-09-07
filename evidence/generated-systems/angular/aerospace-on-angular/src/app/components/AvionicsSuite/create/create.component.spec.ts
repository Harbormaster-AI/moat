
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAvionicsSuiteComponent } from './create.component';
import { AvionicsSuiteService } from '../../../services/AvionicsSuite.service';
import { Router } from '@angular/router';

describe('CreateAvionicsSuiteComponent', () => {
  let component: CreateAvionicsSuiteComponent;
  let fixture: ComponentFixture<CreateAvionicsSuiteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAvionicsSuiteComponent
      ],
      providers: [
        AvionicsSuiteService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAvionicsSuiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});