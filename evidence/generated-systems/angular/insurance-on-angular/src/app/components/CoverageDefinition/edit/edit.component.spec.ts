
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditCoverageDefinitionComponent } from './edit.component';
import { CoverageDefinitionService } from '../../../services/CoverageDefinition.service';

describe('EditCoverageDefinitionComponent', () => {
  let component: EditCoverageDefinitionComponent;
  let fixture: ComponentFixture<EditCoverageDefinitionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditCoverageDefinitionComponent
      ],
      providers: [
        CoverageDefinitionService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditCoverageDefinitionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});