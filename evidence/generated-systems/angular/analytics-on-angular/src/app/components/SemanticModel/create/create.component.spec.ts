
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSemanticModelComponent } from './create.component';
import { SemanticModelService } from '../../../services/SemanticModel.service';
import { Router } from '@angular/router';

describe('CreateSemanticModelComponent', () => {
  let component: CreateSemanticModelComponent;
  let fixture: ComponentFixture<CreateSemanticModelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSemanticModelComponent
      ],
      providers: [
        SemanticModelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSemanticModelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});