
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexNotebookComponent } from './index.component';
import { NotebookService } from '../../../services/Notebook.service';

describe('IndexNotebookComponent', () => {
  let component: IndexNotebookComponent;
  let fixture: ComponentFixture<IndexNotebookComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexNotebookComponent
      ],
      providers: [
        NotebookService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexNotebookComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});