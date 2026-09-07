
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateNotebookComponent } from './create.component';
import { NotebookService } from '../../../services/Notebook.service';
import { Router } from '@angular/router';

describe('CreateNotebookComponent', () => {
  let component: CreateNotebookComponent;
  let fixture: ComponentFixture<CreateNotebookComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateNotebookComponent
      ],
      providers: [
        NotebookService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateNotebookComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});