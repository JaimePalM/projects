# Introduction

## Machine Learning

Machine Learning is a branch of Artificial Intelligence (AI) that focuses on the development of algorithms and models that can learn from data and make predictions or decisions without being explicitly programmed.

It can be divided in three types:
- Supervised Learning: use labelled data to train a model to predict labels for new data.
- Unsupervised Learning: use unlabelled data to discover patterns or relationships in the data, can be challenging.
- Reinforcement Learning: use a reward signal to train a model to make decisions or take actions in an environment. It learns from try and error.

Applications of Machine Learning: 
- Image and Speech Recognition (computer vision)
- Natural Language Processing (chat bots)
- Anomaly Detection
- Predictive Maintenance 
- Social Media Analysis (sentiment analysis)

## Deep Learning

Deep Learning is a subset of Machine Learning that uses neural networks with multiple layers to learn patterns and relationships in data.

For example, we want to recognize animals in a photo. One solution is to use ML, we'll need to collect lots of photos of animals and label them. Then we'll train a model to predict the animal. But this can be very complex and time-consuming. Deep Learning uses deep neural networks to learn from data and make predictions. The input data is passed through multiple layers of neurons, and the output is a prediction. At each layer, the data is processed and transformed to create a representation of the data that is more meaningful for the task at hand, an abstract representation that captures the patterns and relationships in the data. It can learn by adjusting the weights of the connections between the layers.

Applications of Deep Learning:
- Face ID Authentication
- Self-driving cars
- Cancer Detection

# TensorFlow

TensorFlow is an open-source machine learning framework created by Google. It is used for developing, training, and deploying machine learning models. It has a wide range of libraries and tools for building and deploying deep learning models. Some features of TensorFlow include:
- Higher-level APIs for building and training models
- Support DL, NN, and more
- Large and active community
- Easy integration with other libraries and frameworks

Components of TensorFlow:
- Core: low-level library for building and training models. Graphs composed of nodes (operations) and edges (tensors).
- API: high-level APIs for building and training models. E.g. Keras (high-level API with minimal code), TensorFlow Hub (pre-trained models), etc.

## Installation

1. Install Python (`https://www.python.org/downloads/windows/`)
2. Install TensorFlow: `pip install tensorflow`

## Architecture

The key components of TensorFlow are:
- Graph: a collection of operations and tensors that describe the computation to be performed.
- Tensor: a multi-dimensional array of values, similar to a matrix. There are different types of tensors, e.g. float32, int32, etc.
- Operation: a computation that takes one or more tensors as input and produces one or more tensors as output.
- Session: execution environment for running TensorFlow operations.

In a typical ML or DL application, we'll have a graph that represents the computation we want to perform, and a session that executes the graph. This graphs consist of operations that manipulate tensors to produce the final result.

During the execution of the graph, TensorFlow automatically performs backpropagation to compute the gradients of the loss function with respect to the variables of the model. This allows TensorFlow to update the model parameters to minimize the loss function.

## APIs

TensorFlow provides a wide range of APIs for building and training models. Some popular APIs include:
- Keras: high-level API for building and training models It's easy to use and allows developers to quickly create prototypes and experiment with different architectures.
- TensorFlow Lite: for deploying models on mobile devices.
- TensorFlow.js: for running models in the browser. No installation required. Useful for web applications such as games or chat bots.
- TensorFlow Extended (TFX): platform to build and deploy production ML pipelines. It includes tools for data validation, data pre-processing, model training, evaluation, and deployment.

# ML: Supervised Learning

Involves using labelled data to train a model to predict labels for new data. For example, if we want to predict the price of a house based on features, the goal would be to learn a mapping between the features of the house and its price. Supervised learning can be divided in two types:
- Regression: predict a continuous value, such as the price of a house
- Classification: predict a categorical value, such as whether a house is for sale or not

To train a supervised model, the labeled data is split into training and testing sets. The model is trained on the training data, used to get the weights of the model, and then tested on the testing data to evaluate the performance of the model. Testing data is labeled but the model has to predict the labels and so we know if the model is performing well or not.

### Linear Regression

In linear regression, the goal is to find the best line that fits the data (describes the relationship between the data). This is done by minimizing the sum of the squared errors between the predicted and actual values.

For the lab we need to install with `pip` the following packages: `jupyter`, `scipy`, `scikit-learn`, `numpy`, `pandas` and `matplotlib`.

Now we can create a notebook with `jupyter notebook` and start coding.

In summary, linear regression is make a prediction based on the linear relationship between the independent and dependent variables. In other words, obtain the $y = mx + b$ line that best fits the data. Later, $x$ can be substituted with any independent variable and the function will return the predicted value.

### Logistic Regression

Logistic regression is used for classification problems. It uses a sigmoid function to transform the output of the linear model into a probability. So the difference between linear regression and logistic regression is that the latter uses a sigmoid function to transform the output of the linear model into a probability.

In logistic regression the dependent variable is binary, 0 or 1. The output of the model is a probability between 0 and 1. If the probability is greater than 0.5, the output is 1, otherwise 0.

The model parameters are estimated using a method called maximum likelihood estimation (MLE), which finds the parameter that maximize the likelihood of the data given the model.

The performance can be evaluated using metrics such as the Receiver Operating Characteristic (ROC) curve and the Area Under the Curve (AUC).

For the lab, go to `LogisticRegression.ipynb`.

# ML: Unsupervised Learning