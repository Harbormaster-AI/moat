
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCoverageDefinitionComponent } from './create.component';
import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';
import { Router } from '@angular/router';

describe('CreateCoverageDefinitionComponent', () => {
  let component: CreateCoverageDefinitionComponent;
  let fixture: ComponentFixture<CreateCoverageDefinitionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCoverageDefinitionComponent
      ],
      providers: [
        CoverageDefinitionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCoverageDefinitionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});